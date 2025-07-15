// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbolishApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AbolishApiResponseBody
	GetRequestId() *string
}

type AbolishApiResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ016
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbolishApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbolishApiResponseBody) GoString() string {
	return s.String()
}

func (s *AbolishApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbolishApiResponseBody) SetRequestId(v string) *AbolishApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbolishApiResponseBody) Validate() error {
	return dara.Validate(s)
}
