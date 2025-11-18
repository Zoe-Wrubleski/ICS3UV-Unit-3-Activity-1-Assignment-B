/**
 * @author Zoe Wrubleski
 * @version 1.0.0
 * @date 2025-11-18
 * @fileoverview This program calculates the subtotal, tax, tip, total and cost per person of a meal
 */

// set variables
let subtotal: number = 415.50;
let people: number = 8;

// calculate tax of 13%
let tax: number = subtotal * (13/100);

// calculate tip of 10%
let tip: number = subtotal * (10/100);

// calculate total
let total: number = subtotal + tax + tip;

// calculate the cost per person
let perPerson: number = total / people

// display answers
console.log("Subtotal: $" + subtotal);
console.log("Tax (13%): $" + tax);
console.log("Tip (10%): $" + tip);
console.log("Total: $" + total);
console.log("Cost per person: $" + perPerson);

console.log("\nDone");
