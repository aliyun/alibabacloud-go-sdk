// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobErrorCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeJobErrorCodeRequest
	GetClientToken() *string
	SetLanguage(v string) *DescribeJobErrorCodeRequest
	GetLanguage() *string
	SetOwnerId(v string) *DescribeJobErrorCodeRequest
	GetOwnerId() *string
	SetTaskId(v string) *DescribeJobErrorCodeRequest
	GetTaskId() *string
}

type DescribeJobErrorCodeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language of the error message. Valid values:
	//
	// 	- **en*	- (default): English
	//
	// 	- **cn**: Chinese
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OwnerId  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the full backup or restore task.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi0****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeJobErrorCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobErrorCodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobErrorCodeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeJobErrorCodeRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeJobErrorCodeRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeJobErrorCodeRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeJobErrorCodeRequest) SetClientToken(v string) *DescribeJobErrorCodeRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeJobErrorCodeRequest) SetLanguage(v string) *DescribeJobErrorCodeRequest {
	s.Language = &v
	return s
}

func (s *DescribeJobErrorCodeRequest) SetOwnerId(v string) *DescribeJobErrorCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeJobErrorCodeRequest) SetTaskId(v string) *DescribeJobErrorCodeRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeJobErrorCodeRequest) Validate() error {
	return dara.Validate(s)
}
