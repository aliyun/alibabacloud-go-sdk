// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppItemId(v string) *DelAppRequest
	GetAppItemId() *string
}

type DelAppRequest struct {
	// This parameter is required.
	AppItemId *string `json:"AppItemId,omitempty" xml:"AppItemId,omitempty"`
}

func (s DelAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DelAppRequest) GoString() string {
	return s.String()
}

func (s *DelAppRequest) GetAppItemId() *string {
	return s.AppItemId
}

func (s *DelAppRequest) SetAppItemId(v string) *DelAppRequest {
	s.AppItemId = &v
	return s
}

func (s *DelAppRequest) Validate() error {
	return dara.Validate(s)
}
