// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomCallTaggingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomNumberList(v string) *CreateCustomCallTaggingRequest
	GetCustomNumberList() *string
	SetInstanceId(v string) *CreateCustomCallTaggingRequest
	GetInstanceId() *string
}

type CreateCustomCallTaggingRequest struct {
	// A list of inbound control tags, formatted as a JSON array string. The number of array elements must be between 1 and 10,000. Each element in the array is an object with the following properties: number (must be a numeric string of 4 to 32 characters), description, and callTagNameList. The callTagNameList is an array whose elements are number labels (ensure that these number labels have already been created).
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"number":"1312121****","description":"王先生","callTagNameList":["TagA"]},{"number":"1388888****","description":"张先生","callTagNameList":["TagB"]}]
	CustomNumberList *string `json:"CustomNumberList,omitempty" xml:"CustomNumberList,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateCustomCallTaggingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCallTaggingRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomCallTaggingRequest) GetCustomNumberList() *string {
	return s.CustomNumberList
}

func (s *CreateCustomCallTaggingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCustomCallTaggingRequest) SetCustomNumberList(v string) *CreateCustomCallTaggingRequest {
	s.CustomNumberList = &v
	return s
}

func (s *CreateCustomCallTaggingRequest) SetInstanceId(v string) *CreateCustomCallTaggingRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomCallTaggingRequest) Validate() error {
	return dara.Validate(s)
}
