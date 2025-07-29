// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterDiagnosisResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetClusterDiagnosisResultResponseBody
	GetCode() *int64
	SetCreated(v string) *GetClusterDiagnosisResultResponseBody
	GetCreated() *string
	SetDiagnosisId(v string) *GetClusterDiagnosisResultResponseBody
	GetDiagnosisId() *string
	SetFinished(v string) *GetClusterDiagnosisResultResponseBody
	GetFinished() *string
	SetMessage(v string) *GetClusterDiagnosisResultResponseBody
	GetMessage() *string
	SetResult(v string) *GetClusterDiagnosisResultResponseBody
	GetResult() *string
	SetStatus(v int64) *GetClusterDiagnosisResultResponseBody
	GetStatus() *int64
	SetTarget(v string) *GetClusterDiagnosisResultResponseBody
	GetTarget() *string
	SetType(v string) *GetClusterDiagnosisResultResponseBody
	GetType() *string
}

type GetClusterDiagnosisResultResponseBody struct {
	// The code that indicates the diagnostic result. Valid values:
	//
	// 	- 0: the diagnostic is completed.
	//
	// 	- 1: the diagnostic failed.
	//
	// example:
	//
	// 0
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// The time when the diagnostic is initiated.
	//
	// example:
	//
	// 2024-05-28T11:29Z
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The diagnostic ID.
	//
	// example:
	//
	// 6cf6b62e334e4583bdfd26707516****
	DiagnosisId *string `json:"diagnosis_id,omitempty" xml:"diagnosis_id,omitempty"`
	// The time when the diagnostic is completed.
	//
	// example:
	//
	// 2024-05-28T11:29Z
	Finished *string `json:"finished,omitempty" xml:"finished,omitempty"`
	// The diagnostic status information.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The diagnostic result.
	//
	// example:
	//
	// {"phase":5,"version":"20240101"}
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// The status of the diagnostic. Valid values:
	//
	// 	- 0: The diagnostic is created.
	//
	// 	- 1: The diagnostic is running.
	//
	// 	- 2: The diagnostic is completed.
	//
	// example:
	//
	// 2
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// The diagnostic object.
	//
	// example:
	//
	// {"name":"cn-hongkong.10.0.0.246"}
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
	// The type of the diagnostic.
	//
	// Valid values:
	//
	// 	- node
	//
	// 	- ingress
	//
	// 	- cluster
	//
	// 	- memory
	//
	// 	- pod
	//
	// 	- service
	//
	// 	- network
	//
	// example:
	//
	// Node
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetClusterDiagnosisResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterDiagnosisResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterDiagnosisResultResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetClusterDiagnosisResultResponseBody) GetCreated() *string {
	return s.Created
}

func (s *GetClusterDiagnosisResultResponseBody) GetDiagnosisId() *string {
	return s.DiagnosisId
}

func (s *GetClusterDiagnosisResultResponseBody) GetFinished() *string {
	return s.Finished
}

func (s *GetClusterDiagnosisResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetClusterDiagnosisResultResponseBody) GetResult() *string {
	return s.Result
}

func (s *GetClusterDiagnosisResultResponseBody) GetStatus() *int64 {
	return s.Status
}

func (s *GetClusterDiagnosisResultResponseBody) GetTarget() *string {
	return s.Target
}

func (s *GetClusterDiagnosisResultResponseBody) GetType() *string {
	return s.Type
}

func (s *GetClusterDiagnosisResultResponseBody) SetCode(v int64) *GetClusterDiagnosisResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetClusterDiagnosisResultResponseBody) SetCreated(v string) *GetClusterDiagnosisResultResponseBody {
	s.Created = &v
	return s
}

func (s *GetClusterDiagnosisResultResponseBody) SetDiagnosisId(v string) *GetClusterDiagnosisResultResponseBody {
	s.DiagnosisId = &v
	return s
}

func (s *GetClusterDiagnosisResultResponseBody) SetFinished(v string) *GetClusterDiagnosisResultResponseBody {
	s.Finished = &v
	return s
}

func (s *GetClusterDiagnosisResultResponseBody) SetMessage(v string) *GetClusterDiagnosisResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetClusterDiagnosisResultResponseBody) SetResult(v string) *GetClusterDiagnosisResultResponseBody {
	s.Result = &v
	return s
}

func (s *GetClusterDiagnosisResultResponseBody) SetStatus(v int64) *GetClusterDiagnosisResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetClusterDiagnosisResultResponseBody) SetTarget(v string) *GetClusterDiagnosisResultResponseBody {
	s.Target = &v
	return s
}

func (s *GetClusterDiagnosisResultResponseBody) SetType(v string) *GetClusterDiagnosisResultResponseBody {
	s.Type = &v
	return s
}

func (s *GetClusterDiagnosisResultResponseBody) Validate() error {
	return dara.Validate(s)
}
