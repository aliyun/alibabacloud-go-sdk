// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCdnSubTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCdnSubTaskResponseBody
	GetRequestId() *string
}

type CreateCdnSubTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCdnSubTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCdnSubTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCdnSubTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCdnSubTaskResponseBody) SetRequestId(v string) *CreateCdnSubTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCdnSubTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
