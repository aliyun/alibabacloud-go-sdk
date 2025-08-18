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
	// The application ID, which can be obtained by calling the [ListEdgeContainerApps](~~ListEdgeContainerApps~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The number of log entries to output.
	//
	// This parameter is required.
	//
	// example:
	//
	// 500
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
