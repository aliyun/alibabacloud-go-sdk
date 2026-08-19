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
	// 	Notice: This field was of the Long type in SDK versions earlier than 8.0.0 and is of the String type in SDK 8.0.0 and later. **This change does not affect normal SDK usage, and the parameter is still returned in the type defined in the SDK**. The type change may cause project compilation failures only when you upgrade across SDK version 8.0.0. In this case, you must manually correct the data type.
	//
	// example:
	//
	// 631478864897630XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// A5B97987-66EA-5563-9599-A2752292****
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
