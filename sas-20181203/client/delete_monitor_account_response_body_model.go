// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMonitorAccountResponseBody
	GetRequestId() *string
}

type DeleteMonitorAccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMonitorAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMonitorAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMonitorAccountResponseBody) SetRequestId(v string) *DeleteMonitorAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMonitorAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
