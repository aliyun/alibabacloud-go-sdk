// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *DeleteCloudAccessRequest
	GetAccessId() *string
}

type DeleteCloudAccessRequest struct {
	// The ID of the access key.
	//
	// example:
	//
	// 132
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
}

func (s DeleteCloudAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudAccessRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudAccessRequest) GetAccessId() *string {
	return s.AccessId
}

func (s *DeleteCloudAccessRequest) SetAccessId(v string) *DeleteCloudAccessRequest {
	s.AccessId = &v
	return s
}

func (s *DeleteCloudAccessRequest) Validate() error {
	return dara.Validate(s)
}
