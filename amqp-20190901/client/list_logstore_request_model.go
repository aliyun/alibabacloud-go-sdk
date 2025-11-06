// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogstoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *ListLogstoreRequest
	GetConsoleSessionId() *string
	SetProjectName(v string) *ListLogstoreRequest
	GetProjectName() *string
}

type ListLogstoreRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListLogstoreRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLogstoreRequest) GoString() string {
	return s.String()
}

func (s *ListLogstoreRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ListLogstoreRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListLogstoreRequest) SetConsoleSessionId(v string) *ListLogstoreRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ListLogstoreRequest) SetProjectName(v string) *ListLogstoreRequest {
	s.ProjectName = &v
	return s
}

func (s *ListLogstoreRequest) Validate() error {
	return dara.Validate(s)
}
