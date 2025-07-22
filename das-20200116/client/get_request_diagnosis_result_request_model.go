// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRequestDiagnosisResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRequestDiagnosisResultRequest
	GetInstanceId() *string
	SetMessageId(v string) *GetRequestDiagnosisResultRequest
	GetMessageId() *string
	SetNodeId(v string) *GetRequestDiagnosisResultRequest
	GetNodeId() *string
	SetSource(v string) *GetRequestDiagnosisResultRequest
	GetSource() *string
	SetSqlId(v string) *GetRequestDiagnosisResultRequest
	GetSqlId() *string
}

type GetRequestDiagnosisResultRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-0iwhhl8gx0ld6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The unique ID of the diagnostic task.[](~~341609~~)
	//
	// >  If you set MessageId to the task ID of the automatic SQL optimization feature, no result is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61820b594664275c4429****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The node ID.
	//
	// >  You must specify the node ID if your database instance is a PolarDB for MySQL cluster, a PolarDB for PostgreSQL (compatible with Oracle) instance, or an ApsaraDB for MongoDB database.
	//
	// example:
	//
	// 202****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The source of the task.
	//
	// >  This parameter is required if you call this operation in the DAS console. You do not need to specify this parameter when you call this operation.
	//
	// example:
	//
	// None
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The SQL template ID.
	//
	// >  This parameter is required if you call this operation in the DAS console. You do not need to specify this parameter when you call this operation.
	//
	// example:
	//
	// None
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
}

func (s GetRequestDiagnosisResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRequestDiagnosisResultRequest) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRequestDiagnosisResultRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *GetRequestDiagnosisResultRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetRequestDiagnosisResultRequest) GetSource() *string {
	return s.Source
}

func (s *GetRequestDiagnosisResultRequest) GetSqlId() *string {
	return s.SqlId
}

func (s *GetRequestDiagnosisResultRequest) SetInstanceId(v string) *GetRequestDiagnosisResultRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRequestDiagnosisResultRequest) SetMessageId(v string) *GetRequestDiagnosisResultRequest {
	s.MessageId = &v
	return s
}

func (s *GetRequestDiagnosisResultRequest) SetNodeId(v string) *GetRequestDiagnosisResultRequest {
	s.NodeId = &v
	return s
}

func (s *GetRequestDiagnosisResultRequest) SetSource(v string) *GetRequestDiagnosisResultRequest {
	s.Source = &v
	return s
}

func (s *GetRequestDiagnosisResultRequest) SetSqlId(v string) *GetRequestDiagnosisResultRequest {
	s.SqlId = &v
	return s
}

func (s *GetRequestDiagnosisResultRequest) Validate() error {
	return dara.Validate(s)
}
