// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGraphSchemaResponseBody
	GetCode() *string
	SetGraphName(v string) *GetGraphSchemaResponseBody
	GetGraphName() *string
	SetMessage(v string) *GetGraphSchemaResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGraphSchemaResponseBody
	GetRequestId() *string
	SetSchemaVersion(v string) *GetGraphSchemaResponseBody
	GetSchemaVersion() *string
	SetYamlEdit(v string) *GetGraphSchemaResponseBody
	GetYamlEdit() *string
}

type GetGraphSchemaResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 图谱名称
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 当前 active Graph Schema 版本
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	SchemaVersion *string `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	// 按 READ 权限裁剪的 Graph Schema 原始 YAML 文本，保留授权子图内的 $ref
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	YamlEdit *string `json:"yamlEdit,omitempty" xml:"yamlEdit,omitempty"`
}

func (s GetGraphSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGraphSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetGraphSchemaResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGraphSchemaResponseBody) GetGraphName() *string {
	return s.GraphName
}

func (s *GetGraphSchemaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGraphSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGraphSchemaResponseBody) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *GetGraphSchemaResponseBody) GetYamlEdit() *string {
	return s.YamlEdit
}

func (s *GetGraphSchemaResponseBody) SetCode(v string) *GetGraphSchemaResponseBody {
	s.Code = &v
	return s
}

func (s *GetGraphSchemaResponseBody) SetGraphName(v string) *GetGraphSchemaResponseBody {
	s.GraphName = &v
	return s
}

func (s *GetGraphSchemaResponseBody) SetMessage(v string) *GetGraphSchemaResponseBody {
	s.Message = &v
	return s
}

func (s *GetGraphSchemaResponseBody) SetRequestId(v string) *GetGraphSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGraphSchemaResponseBody) SetSchemaVersion(v string) *GetGraphSchemaResponseBody {
	s.SchemaVersion = &v
	return s
}

func (s *GetGraphSchemaResponseBody) SetYamlEdit(v string) *GetGraphSchemaResponseBody {
	s.YamlEdit = &v
	return s
}

func (s *GetGraphSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}
