// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterMiguDownloadSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceId(v string) *ModelRouterMiguDownloadSourceRequest
	GetSourceId() *string
}

type ModelRouterMiguDownloadSourceRequest struct {
	// The unique identifier of the source file. This is the sourceId returned by the upload operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3f2a1b9c8d7e4f60a1b2c3d4e5f6a7b8
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s ModelRouterMiguDownloadSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterMiguDownloadSourceRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterMiguDownloadSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ModelRouterMiguDownloadSourceRequest) SetSourceId(v string) *ModelRouterMiguDownloadSourceRequest {
	s.SourceId = &v
	return s
}

func (s *ModelRouterMiguDownloadSourceRequest) Validate() error {
	return dara.Validate(s)
}
