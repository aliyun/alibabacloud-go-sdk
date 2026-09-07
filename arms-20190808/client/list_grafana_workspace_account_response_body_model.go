// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrafanaWorkspaceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListGrafanaWorkspaceAccountResponseBody
	GetCode() *int32
	SetData(v []*GrafanaWorkspaceAccount) *ListGrafanaWorkspaceAccountResponseBody
	GetData() []*GrafanaWorkspaceAccount
	SetMessage(v string) *ListGrafanaWorkspaceAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGrafanaWorkspaceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGrafanaWorkspaceAccountResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListGrafanaWorkspaceAccountResponseBody
	GetTraceId() *string
}

type ListGrafanaWorkspaceAccountResponseBody struct {
	// example:
	//
	// 200
	Code *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GrafanaWorkspaceAccount `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4789C3E9-A85A-524B-B97B-9D2B14BA06BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// eac0a8048716731735000007137d000b
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListGrafanaWorkspaceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGrafanaWorkspaceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListGrafanaWorkspaceAccountResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListGrafanaWorkspaceAccountResponseBody) GetData() []*GrafanaWorkspaceAccount {
	return s.Data
}

func (s *ListGrafanaWorkspaceAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGrafanaWorkspaceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGrafanaWorkspaceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGrafanaWorkspaceAccountResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListGrafanaWorkspaceAccountResponseBody) SetCode(v int32) *ListGrafanaWorkspaceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *ListGrafanaWorkspaceAccountResponseBody) SetData(v []*GrafanaWorkspaceAccount) *ListGrafanaWorkspaceAccountResponseBody {
	s.Data = v
	return s
}

func (s *ListGrafanaWorkspaceAccountResponseBody) SetMessage(v string) *ListGrafanaWorkspaceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *ListGrafanaWorkspaceAccountResponseBody) SetRequestId(v string) *ListGrafanaWorkspaceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGrafanaWorkspaceAccountResponseBody) SetSuccess(v bool) *ListGrafanaWorkspaceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *ListGrafanaWorkspaceAccountResponseBody) SetTraceId(v string) *ListGrafanaWorkspaceAccountResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListGrafanaWorkspaceAccountResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
