// bitblt.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

//#include <iostream>
#include <windows.h>
#include <vector> // list vs vector?

struct Point
{
    int x;
    int y;
};

int main()
{
    //std::cout << "Hello World!\n";
    printf("hello world\n");
    std::vector<Point> points;

    Point p = { 0 };
    p.x = 100;
    p.y = 10;
    points.push_back(p);
    p.x = 20;
    p.y = 10;
    points.push_back(p);
    for (size_t i = 0; i < points.size(); i++)
    {
        //printf("%d (%d, %d)\n", i, points[i].x, points[i].y);
    }

    BITMAP bm; // TODO dont need this
    HBITMAP bi = (HBITMAP)LoadImageA(0, "dead.bmp", IMAGE_BITMAP, 10, 10, LR_LOADFROMFILE);
    // TODO what if image isnt there
    HDC whdc = GetDC(NULL);
    HDC hdcMem = CreateCompatibleDC(whdc);
    SelectObject(hdcMem, bi);
    GetObject((HGDIOBJ)bi, sizeof(bm), &bm);
    while (1)
    {
        BitBlt(whdc, 100, 10, bm.bmWidth, bm.bmHeight, hdcMem, 0, 0, SRCCOPY);
        BitBlt(whdc, 20, 10, bm.bmWidth, bm.bmHeight, hdcMem, 0, 0, SRCCOPY);
    }
    DeleteDC(hdcMem);
    DeleteObject(bi);
}

// Run program: Ctrl + F5 or Debug > Start Without Debugging menu
// Debug program: F5 or Debug > Start Debugging menu

// Tips for Getting Started: 
//   1. Use the Solution Explorer window to add/manage files
//   2. Use the Team Explorer window to connect to source control
//   3. Use the Output window to see build output and other messages
//   4. Use the Error List window to view errors
//   5. Go to Project > Add New Item to create new code files, or Project > Add Existing Item to add existing code files to the project
//   6. In the future, to open this project again, go to File > Open > Project and select the .sln file
