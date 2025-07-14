// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteJobRequest
	GetAppId() *string
}

type DeleteJobRequest struct {
	// The ID of the job template that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 017f39b8-dfa4-4e16-a84b-1dcee4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteJobRequest) SetAppId(v string) *DeleteJobRequest {
	s.AppId = &v
	return s
}

func (s *DeleteJobRequest) Validate() error {
	return dara.Validate(s)
}
