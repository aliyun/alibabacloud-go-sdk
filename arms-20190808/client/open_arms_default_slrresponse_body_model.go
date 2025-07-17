// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenArmsDefaultSLRResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *OpenArmsDefaultSLRResponseBody
	GetData() *string
	SetRequestId(v string) *OpenArmsDefaultSLRResponseBody
	GetRequestId() *string
}

type OpenArmsDefaultSLRResponseBody struct {
	// Indicates whether the call was successful. Valid values:
	//
	// 	- `true`: The call was successful.
	//
	// 	- `false`: The call failed.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 53CACA70-2CF7-490C-BD06-1A2AE4EB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenArmsDefaultSLRResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenArmsDefaultSLRResponseBody) GoString() string {
	return s.String()
}

func (s *OpenArmsDefaultSLRResponseBody) GetData() *string {
	return s.Data
}

func (s *OpenArmsDefaultSLRResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenArmsDefaultSLRResponseBody) SetData(v string) *OpenArmsDefaultSLRResponseBody {
	s.Data = &v
	return s
}

func (s *OpenArmsDefaultSLRResponseBody) SetRequestId(v string) *OpenArmsDefaultSLRResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenArmsDefaultSLRResponseBody) Validate() error {
	return dara.Validate(s)
}
