// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTempFileTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTempFileTaskIds(v []*string) *DeleteTempFileTasksRequest
	GetTempFileTaskIds() []*string
}

type DeleteTempFileTasksRequest struct {
	TempFileTaskIds []*string `json:"TempFileTaskIds,omitempty" xml:"TempFileTaskIds,omitempty" type:"Repeated"`
}

func (s DeleteTempFileTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTempFileTasksRequest) GoString() string {
	return s.String()
}

func (s *DeleteTempFileTasksRequest) GetTempFileTaskIds() []*string {
	return s.TempFileTaskIds
}

func (s *DeleteTempFileTasksRequest) SetTempFileTaskIds(v []*string) *DeleteTempFileTasksRequest {
	s.TempFileTaskIds = v
	return s
}

func (s *DeleteTempFileTasksRequest) Validate() error {
	return dara.Validate(s)
}
