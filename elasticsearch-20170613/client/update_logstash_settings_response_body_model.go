// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogstashSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLogstashSettingsResponseBody
	GetRequestId() *string
}

type UpdateLogstashSettingsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 71D0D1DA-B22F-58CB-AF5B-D1657A6A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLogstashSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLogstashSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLogstashSettingsResponseBody) SetRequestId(v string) *UpdateLogstashSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLogstashSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
