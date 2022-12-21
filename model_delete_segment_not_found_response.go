/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.0.2
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// DeleteSegmentNotFoundResponse struct for DeleteSegmentNotFoundResponse
type DeleteSegmentNotFoundResponse struct {
	Success *bool `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeleteSegmentNotFoundResponse DeleteSegmentNotFoundResponse

// NewDeleteSegmentNotFoundResponse instantiates a new DeleteSegmentNotFoundResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSegmentNotFoundResponse() *DeleteSegmentNotFoundResponse {
	this := DeleteSegmentNotFoundResponse{}
	return &this
}

// NewDeleteSegmentNotFoundResponseWithDefaults instantiates a new DeleteSegmentNotFoundResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSegmentNotFoundResponseWithDefaults() *DeleteSegmentNotFoundResponse {
	this := DeleteSegmentNotFoundResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DeleteSegmentNotFoundResponse) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSegmentNotFoundResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DeleteSegmentNotFoundResponse) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DeleteSegmentNotFoundResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o DeleteSegmentNotFoundResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeleteSegmentNotFoundResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDeleteSegmentNotFoundResponse := _DeleteSegmentNotFoundResponse{}

	if err = json.Unmarshal(bytes, &varDeleteSegmentNotFoundResponse); err == nil {
		*o = DeleteSegmentNotFoundResponse(varDeleteSegmentNotFoundResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteSegmentNotFoundResponse struct {
	value *DeleteSegmentNotFoundResponse
	isSet bool
}

func (v NullableDeleteSegmentNotFoundResponse) Get() *DeleteSegmentNotFoundResponse {
	return v.value
}

func (v *NullableDeleteSegmentNotFoundResponse) Set(val *DeleteSegmentNotFoundResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSegmentNotFoundResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSegmentNotFoundResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSegmentNotFoundResponse(val *DeleteSegmentNotFoundResponse) *NullableDeleteSegmentNotFoundResponse {
	return &NullableDeleteSegmentNotFoundResponse{value: val, isSet: true}
}

func (v NullableDeleteSegmentNotFoundResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSegmentNotFoundResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


