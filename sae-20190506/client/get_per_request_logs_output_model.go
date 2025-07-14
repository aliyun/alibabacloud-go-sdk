// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPerRequestLogsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetPerRequestLogsOutput
	GetData() *string
	SetRequestId(v string) *GetPerRequestLogsOutput
	GetRequestId() *string
}

type GetPerRequestLogsOutput struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPerRequestLogsOutput) String() string {
	return dara.Prettify(s)
}

func (s GetPerRequestLogsOutput) GoString() string {
	return s.String()
}

func (s *GetPerRequestLogsOutput) GetData() *string {
	return s.Data
}

func (s *GetPerRequestLogsOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPerRequestLogsOutput) SetData(v string) *GetPerRequestLogsOutput {
	s.Data = &v
	return s
}

func (s *GetPerRequestLogsOutput) SetRequestId(v string) *GetPerRequestLogsOutput {
	s.RequestId = &v
	return s
}

func (s *GetPerRequestLogsOutput) Validate() error {
	return dara.Validate(s)
}
