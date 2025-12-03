// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineEmasArtifactUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceConnectionId(v int64) *GetPipelineEmasArtifactUrlRequest
	GetServiceConnectionId() *int64
}

type GetPipelineEmasArtifactUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 122
	ServiceConnectionId *int64 `json:"serviceConnectionId,omitempty" xml:"serviceConnectionId,omitempty"`
}

func (s GetPipelineEmasArtifactUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineEmasArtifactUrlRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineEmasArtifactUrlRequest) GetServiceConnectionId() *int64 {
	return s.ServiceConnectionId
}

func (s *GetPipelineEmasArtifactUrlRequest) SetServiceConnectionId(v int64) *GetPipelineEmasArtifactUrlRequest {
	s.ServiceConnectionId = &v
	return s
}

func (s *GetPipelineEmasArtifactUrlRequest) Validate() error {
	return dara.Validate(s)
}
