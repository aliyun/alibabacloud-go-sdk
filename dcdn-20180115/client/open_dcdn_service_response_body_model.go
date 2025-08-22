// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDcdnServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenDcdnServiceResponseBody
	GetRequestId() *string
}

type OpenDcdnServiceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 97C68796-EB7F-4D41-9D5B-12B909D76503
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenDcdnServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenDcdnServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenDcdnServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenDcdnServiceResponseBody) SetRequestId(v string) *OpenDcdnServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenDcdnServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
