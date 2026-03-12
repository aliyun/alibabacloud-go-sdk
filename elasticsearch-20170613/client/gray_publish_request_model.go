// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrayPublishRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXRequestChangeId(v string) *GrayPublishRequest
	GetXRequestChangeId() *string
}

type GrayPublishRequest struct {
	// example:
	//
	// ea8d33aa4371c3499d0***
	XRequestChangeId *string `json:"X-Request-ChangeId,omitempty" xml:"X-Request-ChangeId,omitempty"`
}

func (s GrayPublishRequest) String() string {
	return dara.Prettify(s)
}

func (s GrayPublishRequest) GoString() string {
	return s.String()
}

func (s *GrayPublishRequest) GetXRequestChangeId() *string {
	return s.XRequestChangeId
}

func (s *GrayPublishRequest) SetXRequestChangeId(v string) *GrayPublishRequest {
	s.XRequestChangeId = &v
	return s
}

func (s *GrayPublishRequest) Validate() error {
	return dara.Validate(s)
}
