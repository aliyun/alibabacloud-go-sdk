// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOcspStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyOcspStatusResponseBody
	GetRequestId() *string
}

type ModifyOcspStatusResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D8DDBA8E-8182-5C85-AA41-F17EACFCAE0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOcspStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOcspStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOcspStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOcspStatusResponseBody) SetRequestId(v string) *ModifyOcspStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOcspStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
