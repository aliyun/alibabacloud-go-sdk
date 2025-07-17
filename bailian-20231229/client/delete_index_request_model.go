// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexId(v string) *DeleteIndexRequest
	GetIndexId() *string
}

type DeleteIndexRequest struct {
	// The primary key ID of the knowledge base, which is the `Data.Id` parameter returned by the [CreateIndex](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-createindex) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// f89ie0fj5a
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
}

func (s DeleteIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexRequest) GoString() string {
	return s.String()
}

func (s *DeleteIndexRequest) GetIndexId() *string {
	return s.IndexId
}

func (s *DeleteIndexRequest) SetIndexId(v string) *DeleteIndexRequest {
	s.IndexId = &v
	return s
}

func (s *DeleteIndexRequest) Validate() error {
	return dara.Validate(s)
}
