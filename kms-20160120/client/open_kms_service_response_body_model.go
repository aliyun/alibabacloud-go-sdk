// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenKmsServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenKmsServiceResponseBody
	GetRequestId() *string
}

type OpenKmsServiceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3455b9b4-95c1-419d-b310-db6a53b09a39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenKmsServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenKmsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenKmsServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenKmsServiceResponseBody) SetRequestId(v string) *OpenKmsServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenKmsServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
