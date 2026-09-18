// Build: 35dd22a46bc0db388a9e1e6fc6049df5
using System;

internal static class Utilities
{
    public static int Clamp(int value, int minimum, int maximum)
        => Math.Min(maximum, Math.Max(minimum, value));
}
