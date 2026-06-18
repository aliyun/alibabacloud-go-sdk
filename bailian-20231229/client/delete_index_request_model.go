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
	// The knowledge base ID, which is the `Data.Id` returned by the **CreateIndex*	- operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// f89ie0xxxx
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
