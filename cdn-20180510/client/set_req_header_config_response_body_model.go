// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetReqHeaderConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetReqHeaderConfigResponseBody
	GetRequestId() *string
}

type SetReqHeaderConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetReqHeaderConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetReqHeaderConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetReqHeaderConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetReqHeaderConfigResponseBody) SetRequestId(v string) *SetReqHeaderConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetReqHeaderConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
