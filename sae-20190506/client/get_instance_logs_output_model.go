// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLogsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetInstanceLogsOutput
	GetData() *string
	SetRequestId(v string) *GetInstanceLogsOutput
	GetRequestId() *string
}

type GetInstanceLogsOutput struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetInstanceLogsOutput) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLogsOutput) GoString() string {
	return s.String()
}

func (s *GetInstanceLogsOutput) GetData() *string {
	return s.Data
}

func (s *GetInstanceLogsOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceLogsOutput) SetData(v string) *GetInstanceLogsOutput {
	s.Data = &v
	return s
}

func (s *GetInstanceLogsOutput) SetRequestId(v string) *GetInstanceLogsOutput {
	s.RequestId = &v
	return s
}

func (s *GetInstanceLogsOutput) Validate() error {
	return dara.Validate(s)
}
