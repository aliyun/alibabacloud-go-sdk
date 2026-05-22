// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetEdgeContainerLogsRequest
	GetAppId() *string
	SetLines(v int32) *GetEdgeContainerLogsRequest
	GetLines() *int32
}

type GetEdgeContainerLogsRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Lines *int32 `json:"Lines,omitempty" xml:"Lines,omitempty"`
}

func (s GetEdgeContainerLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerLogsRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerLogsRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetEdgeContainerLogsRequest) GetLines() *int32 {
	return s.Lines
}

func (s *GetEdgeContainerLogsRequest) SetAppId(v string) *GetEdgeContainerLogsRequest {
	s.AppId = &v
	return s
}

func (s *GetEdgeContainerLogsRequest) SetLines(v int32) *GetEdgeContainerLogsRequest {
	s.Lines = &v
	return s
}

func (s *GetEdgeContainerLogsRequest) Validate() error {
	return dara.Validate(s)
}
