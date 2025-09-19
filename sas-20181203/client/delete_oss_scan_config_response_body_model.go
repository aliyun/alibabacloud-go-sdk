// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOssScanConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOssScanConfigResponseBody
	GetRequestId() *string
}

type DeleteOssScanConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB393****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOssScanConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOssScanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOssScanConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOssScanConfigResponseBody) SetRequestId(v string) *DeleteOssScanConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOssScanConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
