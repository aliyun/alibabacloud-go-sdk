// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncSourceSelectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetContextDatabaseSyncSourceSelectionResponseBody
	GetCode() *string
	SetData(v *GetContextDatabaseSyncSourceSelectionResponseBodyData) *GetContextDatabaseSyncSourceSelectionResponseBody
	GetData() *GetContextDatabaseSyncSourceSelectionResponseBodyData
	SetHttpStatusCode(v int64) *GetContextDatabaseSyncSourceSelectionResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetContextDatabaseSyncSourceSelectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetContextDatabaseSyncSourceSelectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetContextDatabaseSyncSourceSelectionResponseBody
	GetSuccess() *bool
}

type GetContextDatabaseSyncSourceSelectionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Sample value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetContextDatabaseSyncSourceSelectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message.
	//
	// example:
	//
	// Sample value
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// id-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetContextDatabaseSyncSourceSelectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncSourceSelectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBody) GetData() *GetContextDatabaseSyncSourceSelectionResponseBodyData {
	return s.Data
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBody) SetCode(v string) *GetContextDatabaseSyncSourceSelectionResponseBody {
	s.Code = &v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBody) SetData(v *GetContextDatabaseSyncSourceSelectionResponseBodyData) *GetContextDatabaseSyncSourceSelectionResponseBody {
	s.Data = v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBody) SetHttpStatusCode(v int64) *GetContextDatabaseSyncSourceSelectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBody) SetMessage(v string) *GetContextDatabaseSyncSourceSelectionResponseBody {
	s.Message = &v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBody) SetRequestId(v string) *GetContextDatabaseSyncSourceSelectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBody) SetSuccess(v bool) *GetContextDatabaseSyncSourceSelectionResponseBody {
	s.Success = &v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextDatabaseSyncSourceSelectionResponseBodyData struct {
	// The confirmation time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	ConfirmedAt *string `json:"ConfirmedAt,omitempty" xml:"ConfirmedAt,omitempty"`
	// The list of selected node identifiers.
	//
	// example:
	//
	// ["doc:1234567890"]
	SelectedNodeKeys []*string `json:"SelectedNodeKeys,omitempty" xml:"SelectedNodeKeys,omitempty" type:"Repeated"`
	// The selection policy.
	//
	// example:
	//
	// subtree
	SelectionPolicy *string `json:"SelectionPolicy,omitempty" xml:"SelectionPolicy,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s GetContextDatabaseSyncSourceSelectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncSourceSelectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBodyData) GetConfirmedAt() *string {
	return s.ConfirmedAt
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBodyData) GetSelectedNodeKeys() []*string {
	return s.SelectedNodeKeys
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBodyData) GetSelectionPolicy() *string {
	return s.SelectionPolicy
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBodyData) SetConfirmedAt(v string) *GetContextDatabaseSyncSourceSelectionResponseBodyData {
	s.ConfirmedAt = &v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBodyData) SetSelectedNodeKeys(v []*string) *GetContextDatabaseSyncSourceSelectionResponseBodyData {
	s.SelectedNodeKeys = v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBodyData) SetSelectionPolicy(v string) *GetContextDatabaseSyncSourceSelectionResponseBodyData {
	s.SelectionPolicy = &v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBodyData) SetUpdatedAt(v string) *GetContextDatabaseSyncSourceSelectionResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
