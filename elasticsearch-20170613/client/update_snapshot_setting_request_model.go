// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSnapshotSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdateSnapshotSettingRequest
	GetBody() *string
}

type UpdateSnapshotSettingRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSnapshotSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSnapshotSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotSettingRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateSnapshotSettingRequest) SetBody(v string) *UpdateSnapshotSettingRequest {
	s.Body = &v
	return s
}

func (s *UpdateSnapshotSettingRequest) Validate() error {
	return dara.Validate(s)
}
