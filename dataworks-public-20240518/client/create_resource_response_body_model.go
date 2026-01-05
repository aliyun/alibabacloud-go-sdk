// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateResourceResponseBody
	GetId() *string
	SetRequestId(v string) *CreateResourceResponseBody
	GetRequestId() *string
}

type CreateResourceResponseBody struct {
	// The unique identifier of the resource file.
	//
	// > This field is of type Long in SDK versions prior to 8.0.0, and of type String in SDK version 8.0.0 and later. This change does not affect the normal use of the SDK; parameters are still returned according to the type defined in the SDK. Compilation failures due to the type change may occur only when upgrading the SDK across version 8.0.0, in which case users need to manually correct the data type.
	//
	// example:
	//
	// 631478864897630XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// A5B97987-66EA-5563-9599-A2752292XXXX
	//
	// example:
	//
	// The ID of the file resource.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceResponseBody) SetId(v string) *CreateResourceResponseBody {
	s.Id = &v
	return s
}

func (s *CreateResourceResponseBody) SetRequestId(v string) *CreateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
