// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteCloudAppRequest
	GetAppId() *string
}

type DeleteCloudAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteCloudAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteCloudAppRequest) SetAppId(v string) *DeleteCloudAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteCloudAppRequest) Validate() error {
	return dara.Validate(s)
}
