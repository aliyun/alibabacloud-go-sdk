// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleWorkforce interface {
	dara.Model
	String() string
	GoString() string
	SetUserIds(v []*int64) *SimpleWorkforce
	GetUserIds() []*int64
	SetWorkNodeId(v int32) *SimpleWorkforce
	GetWorkNodeId() *int32
}

type SimpleWorkforce struct {
	// List of personnel information.
	UserIds []*int64 `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
	// Node ID. For more information, see [GetTaskWorkforce](https://help.aliyun.com/document_detail/454697.html).
	//
	// example:
	//
	// 2
	WorkNodeId *int32 `json:"WorkNodeId,omitempty" xml:"WorkNodeId,omitempty"`
}

func (s SimpleWorkforce) String() string {
	return dara.Prettify(s)
}

func (s SimpleWorkforce) GoString() string {
	return s.String()
}

func (s *SimpleWorkforce) GetUserIds() []*int64 {
	return s.UserIds
}

func (s *SimpleWorkforce) GetWorkNodeId() *int32 {
	return s.WorkNodeId
}

func (s *SimpleWorkforce) SetUserIds(v []*int64) *SimpleWorkforce {
	s.UserIds = v
	return s
}

func (s *SimpleWorkforce) SetWorkNodeId(v int32) *SimpleWorkforce {
	s.WorkNodeId = &v
	return s
}

func (s *SimpleWorkforce) Validate() error {
	return dara.Validate(s)
}
