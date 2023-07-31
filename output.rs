use std::ffi::c_char;
use std::ffi::CString;
use std::ffi::CStr;

#[no_mangle]
pub extern "C" fn process_data(input: *const c_char) -> *mut c_char {
    // let output = input.clone();

    let c_str: &CStr = unsafe { CStr::from_ptr(input) };
   
    let output = c_str.to_str().unwrap().to_uppercase();


    CString::new(output).unwrap().into_raw()
}