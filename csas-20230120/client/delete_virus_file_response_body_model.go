// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirusFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVirusFileResponseBody
	GetRequestId() *string
}

type DeleteVirusFileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirusFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirusFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirusFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVirusFileResponseBody) SetRequestId(v string) *DeleteVirusFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVirusFileResponseBody) Validate() error {
	return dara.Validate(s)
}
