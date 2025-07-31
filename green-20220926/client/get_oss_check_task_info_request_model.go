// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckTaskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParentTaskId(v string) *GetOssCheckTaskInfoRequest
	GetParentTaskId() *string
}

type GetOssCheckTaskInfoRequest struct {
	// example:
	//
	// P_AAA**
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
}

func (s GetOssCheckTaskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *GetOssCheckTaskInfoRequest) GetParentTaskId() *string {
	return s.ParentTaskId
}

func (s *GetOssCheckTaskInfoRequest) SetParentTaskId(v string) *GetOssCheckTaskInfoRequest {
	s.ParentTaskId = &v
	return s
}

func (s *GetOssCheckTaskInfoRequest) Validate() error {
	return dara.Validate(s)
}
