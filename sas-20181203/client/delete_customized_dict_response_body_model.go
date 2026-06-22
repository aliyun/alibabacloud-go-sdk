// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizedDictResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomizedDictResponseBody
	GetRequestId() *string
}

type DeleteCustomizedDictResponseBody struct {
	// The ID of the request. The ID is a unique identifier that Alibaba Cloud generates for the request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 85B5D55F-B341-528F-A2CA-AB1207F9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomizedDictResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizedDictResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomizedDictResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomizedDictResponseBody) SetRequestId(v string) *DeleteCustomizedDictResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomizedDictResponseBody) Validate() error {
	return dara.Validate(s)
}
