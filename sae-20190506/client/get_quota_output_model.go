// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaOutput interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetQuotaOutput
	GetRequestId() *string
	SetInstanceLimit(v int64) *GetQuotaOutput
	GetInstanceLimit() *int64
	SetInstanceUsed(v int64) *GetQuotaOutput
	GetInstanceUsed() *int64
}

type GetQuotaOutput struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceLimit *int64  `json:"instanceLimit,omitempty" xml:"instanceLimit,omitempty"`
	InstanceUsed  *int64  `json:"instanceUsed,omitempty" xml:"instanceUsed,omitempty"`
}

func (s GetQuotaOutput) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaOutput) GoString() string {
	return s.String()
}

func (s *GetQuotaOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQuotaOutput) GetInstanceLimit() *int64 {
	return s.InstanceLimit
}

func (s *GetQuotaOutput) GetInstanceUsed() *int64 {
	return s.InstanceUsed
}

func (s *GetQuotaOutput) SetRequestId(v string) *GetQuotaOutput {
	s.RequestId = &v
	return s
}

func (s *GetQuotaOutput) SetInstanceLimit(v int64) *GetQuotaOutput {
	s.InstanceLimit = &v
	return s
}

func (s *GetQuotaOutput) SetInstanceUsed(v int64) *GetQuotaOutput {
	s.InstanceUsed = &v
	return s
}

func (s *GetQuotaOutput) Validate() error {
	return dara.Validate(s)
}
