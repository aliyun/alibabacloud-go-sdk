// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSlsOpenStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSlsOpenStatusResponseBody
	GetRequestId() *string
	SetSlsOpenStatus(v bool) *GetSlsOpenStatusResponseBody
	GetSlsOpenStatus() *bool
}

type GetSlsOpenStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D01666F5-541B-4C78-98A6-D29E02DAAC7C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether Simple Log Service was activated. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	SlsOpenStatus *bool `json:"SlsOpenStatus,omitempty" xml:"SlsOpenStatus,omitempty"`
}

func (s GetSlsOpenStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSlsOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSlsOpenStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSlsOpenStatusResponseBody) GetSlsOpenStatus() *bool {
	return s.SlsOpenStatus
}

func (s *GetSlsOpenStatusResponseBody) SetRequestId(v string) *GetSlsOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSlsOpenStatusResponseBody) SetSlsOpenStatus(v bool) *GetSlsOpenStatusResponseBody {
	s.SlsOpenStatus = &v
	return s
}

func (s *GetSlsOpenStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
