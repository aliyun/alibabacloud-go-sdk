// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceObjectBindingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReplaceObjectBindingsResponseBody
	GetCode() *string
	SetMessage(v string) *ReplaceObjectBindingsResponseBody
	GetMessage() *string
	SetObjectBindings(v []*ReplaceObjectBindingsResponseBodyObjectBindings) *ReplaceObjectBindingsResponseBody
	GetObjectBindings() []*ReplaceObjectBindingsResponseBodyObjectBindings
	SetRequestId(v string) *ReplaceObjectBindingsResponseBody
	GetRequestId() *string
	SetSourceId(v string) *ReplaceObjectBindingsResponseBody
	GetSourceId() *string
}

type ReplaceObjectBindingsResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述，成功时为空
	Message        *string                                            `json:"message,omitempty" xml:"message,omitempty"`
	ObjectBindings []*ReplaceObjectBindingsResponseBodyObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 数据源 ID
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s ReplaceObjectBindingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceObjectBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceObjectBindingsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReplaceObjectBindingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReplaceObjectBindingsResponseBody) GetObjectBindings() []*ReplaceObjectBindingsResponseBodyObjectBindings {
	return s.ObjectBindings
}

func (s *ReplaceObjectBindingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceObjectBindingsResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *ReplaceObjectBindingsResponseBody) SetCode(v string) *ReplaceObjectBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *ReplaceObjectBindingsResponseBody) SetMessage(v string) *ReplaceObjectBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *ReplaceObjectBindingsResponseBody) SetObjectBindings(v []*ReplaceObjectBindingsResponseBodyObjectBindings) *ReplaceObjectBindingsResponseBody {
	s.ObjectBindings = v
	return s
}

func (s *ReplaceObjectBindingsResponseBody) SetRequestId(v string) *ReplaceObjectBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceObjectBindingsResponseBody) SetSourceId(v string) *ReplaceObjectBindingsResponseBody {
	s.SourceId = &v
	return s
}

func (s *ReplaceObjectBindingsResponseBody) Validate() error {
	if s.ObjectBindings != nil {
		for _, item := range s.ObjectBindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReplaceObjectBindingsResponseBodyObjectBindings struct {
	// 绑定对象归属的语义图谱名（object_id 在该 graph 下唯一，必填）
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// 绑定对象 ID
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// 绑定对象类型（如 customer / project）
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s ReplaceObjectBindingsResponseBodyObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s ReplaceObjectBindingsResponseBodyObjectBindings) GoString() string {
	return s.String()
}

func (s *ReplaceObjectBindingsResponseBodyObjectBindings) GetGraphName() *string {
	return s.GraphName
}

func (s *ReplaceObjectBindingsResponseBodyObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *ReplaceObjectBindingsResponseBodyObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *ReplaceObjectBindingsResponseBodyObjectBindings) SetGraphName(v string) *ReplaceObjectBindingsResponseBodyObjectBindings {
	s.GraphName = &v
	return s
}

func (s *ReplaceObjectBindingsResponseBodyObjectBindings) SetObjectId(v string) *ReplaceObjectBindingsResponseBodyObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *ReplaceObjectBindingsResponseBodyObjectBindings) SetObjectType(v string) *ReplaceObjectBindingsResponseBodyObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *ReplaceObjectBindingsResponseBodyObjectBindings) Validate() error {
	return dara.Validate(s)
}
