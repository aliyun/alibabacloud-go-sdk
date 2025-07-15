// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *CreateApiResponseBody
	GetApiId() *string
	SetRequestId(v string) *CreateApiResponseBody
	GetRequestId() *string
}

type CreateApiResponseBody struct {
	// The ID of the API.
	//
	// example:
	//
	// 8afff6c8c4c6447abb035812e4d66b65
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6C87A26A-6A18-4B8E-8099-705278381A2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiResponseBody) GetApiId() *string {
	return s.ApiId
}

func (s *CreateApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApiResponseBody) SetApiId(v string) *CreateApiResponseBody {
	s.ApiId = &v
	return s
}

func (s *CreateApiResponseBody) SetRequestId(v string) *CreateApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApiResponseBody) Validate() error {
	return dara.Validate(s)
}
