// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIndexJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexId(v string) *SubmitIndexJobRequest
	GetIndexId() *string
}

type SubmitIndexJobRequest struct {
	// The primary key ID of the knowledge base, which is the `Data.Id` parameter returned by the [CreateIndex](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-createindex) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
}

func (s SubmitIndexJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitIndexJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitIndexJobRequest) GetIndexId() *string {
	return s.IndexId
}

func (s *SubmitIndexJobRequest) SetIndexId(v string) *SubmitIndexJobRequest {
	s.IndexId = &v
	return s
}

func (s *SubmitIndexJobRequest) Validate() error {
	return dara.Validate(s)
}
