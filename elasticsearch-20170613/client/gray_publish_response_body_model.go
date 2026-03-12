// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrayPublishResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GrayPublishResponseBody
	GetRequestId() *string
	SetResult(v string) *GrayPublishResponseBody
	GetResult() *string
}

type GrayPublishResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GrayPublishResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrayPublishResponseBody) GoString() string {
	return s.String()
}

func (s *GrayPublishResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrayPublishResponseBody) GetResult() *string {
	return s.Result
}

func (s *GrayPublishResponseBody) SetRequestId(v string) *GrayPublishResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrayPublishResponseBody) SetResult(v string) *GrayPublishResponseBody {
	s.Result = &v
	return s
}

func (s *GrayPublishResponseBody) Validate() error {
	return dara.Validate(s)
}
