// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ModifyFileRequest
	GetContent() *string
	SetPartition(v int32) *ModifyFileRequest
	GetPartition() *int32
	SetFileName(v string) *ModifyFileRequest
	GetFileName() *string
}

type ModifyFileRequest struct {
	// The file content.
	//
	// example:
	//
	// ""
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The number of shards.
	//
	// example:
	//
	// ds=20220713
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	// The name of the file in the full path
	//
	// This parameter is required.
	//
	// example:
	//
	// /schemas/generation_schema.json
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s ModifyFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFileRequest) GoString() string {
	return s.String()
}

func (s *ModifyFileRequest) GetContent() *string {
	return s.Content
}

func (s *ModifyFileRequest) GetPartition() *int32 {
	return s.Partition
}

func (s *ModifyFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *ModifyFileRequest) SetContent(v string) *ModifyFileRequest {
	s.Content = &v
	return s
}

func (s *ModifyFileRequest) SetPartition(v int32) *ModifyFileRequest {
	s.Partition = &v
	return s
}

func (s *ModifyFileRequest) SetFileName(v string) *ModifyFileRequest {
	s.FileName = &v
	return s
}

func (s *ModifyFileRequest) Validate() error {
	return dara.Validate(s)
}
