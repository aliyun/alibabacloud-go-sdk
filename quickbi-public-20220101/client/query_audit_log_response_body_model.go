// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuditLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryAuditLogResponseBody
	GetRequestId() *string
	SetResult(v []*QueryAuditLogResponseBodyResult) *QueryAuditLogResponseBody
	GetResult() []*QueryAuditLogResponseBodyResult
	SetSuccess(v bool) *QueryAuditLogResponseBody
	GetSuccess() *bool
}

type QueryAuditLogResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 78C1AA2D-9201-599E-A0BA-6FC462E57A95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array of log objects.
	Result []*QueryAuditLogResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAuditLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAuditLogResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAuditLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAuditLogResponseBody) GetResult() []*QueryAuditLogResponseBodyResult {
	return s.Result
}

func (s *QueryAuditLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAuditLogResponseBody) SetRequestId(v string) *QueryAuditLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAuditLogResponseBody) SetResult(v []*QueryAuditLogResponseBodyResult) *QueryAuditLogResponseBody {
	s.Result = v
	return s
}

func (s *QueryAuditLogResponseBody) SetSuccess(v bool) *QueryAuditLogResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAuditLogResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAuditLogResponseBodyResult struct {
	// The time when the log was generated.
	//
	// example:
	//
	// 2024-04-16 13:17:39
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The account name of the operator.
	//
	// example:
	//
	// wukaibis
	OperatorAccountName *string `json:"OperatorAccountName,omitempty" xml:"OperatorAccountName,omitempty"`
	// The nickname of the operator.
	//
	// example:
	//
	// buc_344078
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The operation type.
	//
	// example:
	//
	// CREATE
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// The ID of the operation target.
	//
	// example:
	//
	// 1113***************8500
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the operation target.
	//
	// example:
	//
	// test
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the operation target.
	//
	// example:
	//
	// USER
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 87c6b145-090c-43e1-9426-8f93be23****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryAuditLogResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryAuditLogResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAuditLogResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryAuditLogResponseBodyResult) GetOperatorAccountName() *string {
	return s.OperatorAccountName
}

func (s *QueryAuditLogResponseBodyResult) GetOperatorName() *string {
	return s.OperatorName
}

func (s *QueryAuditLogResponseBodyResult) GetOperatorType() *string {
	return s.OperatorType
}

func (s *QueryAuditLogResponseBodyResult) GetTargetId() *string {
	return s.TargetId
}

func (s *QueryAuditLogResponseBodyResult) GetTargetName() *string {
	return s.TargetName
}

func (s *QueryAuditLogResponseBodyResult) GetTargetType() *string {
	return s.TargetType
}

func (s *QueryAuditLogResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryAuditLogResponseBodyResult) SetGmtCreate(v string) *QueryAuditLogResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *QueryAuditLogResponseBodyResult) SetOperatorAccountName(v string) *QueryAuditLogResponseBodyResult {
	s.OperatorAccountName = &v
	return s
}

func (s *QueryAuditLogResponseBodyResult) SetOperatorName(v string) *QueryAuditLogResponseBodyResult {
	s.OperatorName = &v
	return s
}

func (s *QueryAuditLogResponseBodyResult) SetOperatorType(v string) *QueryAuditLogResponseBodyResult {
	s.OperatorType = &v
	return s
}

func (s *QueryAuditLogResponseBodyResult) SetTargetId(v string) *QueryAuditLogResponseBodyResult {
	s.TargetId = &v
	return s
}

func (s *QueryAuditLogResponseBodyResult) SetTargetName(v string) *QueryAuditLogResponseBodyResult {
	s.TargetName = &v
	return s
}

func (s *QueryAuditLogResponseBodyResult) SetTargetType(v string) *QueryAuditLogResponseBodyResult {
	s.TargetType = &v
	return s
}

func (s *QueryAuditLogResponseBodyResult) SetWorkspaceId(v string) *QueryAuditLogResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *QueryAuditLogResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
