// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitStagingRoutineCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *CommitStagingRoutineCodeResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *CommitStagingRoutineCodeResponseBody
	GetRequestId() *string
}

type CommitStagingRoutineCodeResponseBody struct {
	// The version number of the code.
	//
	// example:
	//
	// 1620876959997924701
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5CC228B4-7A67-4016-9C9F-4A4133494A91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CommitStagingRoutineCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CommitStagingRoutineCodeResponseBody) GoString() string {
	return s.String()
}

func (s *CommitStagingRoutineCodeResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *CommitStagingRoutineCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CommitStagingRoutineCodeResponseBody) SetContent(v map[string]interface{}) *CommitStagingRoutineCodeResponseBody {
	s.Content = v
	return s
}

func (s *CommitStagingRoutineCodeResponseBody) SetRequestId(v string) *CommitStagingRoutineCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CommitStagingRoutineCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
