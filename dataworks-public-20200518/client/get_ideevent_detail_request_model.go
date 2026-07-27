// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIDEEventDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *GetIDEEventDetailRequest
	GetMessageId() *string
	SetProjectId(v int64) *GetIDEEventDetailRequest
	GetProjectId() *int64
}

type GetIDEEventDetailRequest struct {
	// The message ID in DataWorks OpenEvent. You can obtain the ID from a received message when an extension point event is triggered.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8abcb91f-d266-4073-b907-2ed67****1
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The DataWorks workspace ID. You can obtain the ID from the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetIDEEventDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIDEEventDetailRequest) GoString() string {
	return s.String()
}

func (s *GetIDEEventDetailRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *GetIDEEventDetailRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetIDEEventDetailRequest) SetMessageId(v string) *GetIDEEventDetailRequest {
	s.MessageId = &v
	return s
}

func (s *GetIDEEventDetailRequest) SetProjectId(v int64) *GetIDEEventDetailRequest {
	s.ProjectId = &v
	return s
}

func (s *GetIDEEventDetailRequest) Validate() error {
	return dara.Validate(s)
}
