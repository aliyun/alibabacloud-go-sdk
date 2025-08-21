// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigLayer4RemarkResponseBody
	GetRequestId() *string
}

type ConfigLayer4RemarkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6E46CC51-36BE-1100-B14C-DAF8381B8F73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigLayer4RemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigLayer4RemarkResponseBody) SetRequestId(v string) *ConfigLayer4RemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigLayer4RemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
