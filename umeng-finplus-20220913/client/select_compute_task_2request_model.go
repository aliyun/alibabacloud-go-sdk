// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectComputeTask2Request interface {
	dara.Model
	String() string
	GoString() string
	SetBcConfirm(v bool) *SelectComputeTask2Request
	GetBcConfirm() *bool
	SetBcId(v int64) *SelectComputeTask2Request
	GetBcId() *int64
	SetClientId(v int64) *SelectComputeTask2Request
	GetClientId() *int64
}

type SelectComputeTask2Request struct {
	BcConfirm *bool  `json:"bcConfirm,omitempty" xml:"bcConfirm,omitempty"`
	BcId      *int64 `json:"bcId,omitempty" xml:"bcId,omitempty"`
	ClientId  *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
}

func (s SelectComputeTask2Request) String() string {
	return dara.Prettify(s)
}

func (s SelectComputeTask2Request) GoString() string {
	return s.String()
}

func (s *SelectComputeTask2Request) GetBcConfirm() *bool {
	return s.BcConfirm
}

func (s *SelectComputeTask2Request) GetBcId() *int64 {
	return s.BcId
}

func (s *SelectComputeTask2Request) GetClientId() *int64 {
	return s.ClientId
}

func (s *SelectComputeTask2Request) SetBcConfirm(v bool) *SelectComputeTask2Request {
	s.BcConfirm = &v
	return s
}

func (s *SelectComputeTask2Request) SetBcId(v int64) *SelectComputeTask2Request {
	s.BcId = &v
	return s
}

func (s *SelectComputeTask2Request) SetClientId(v int64) *SelectComputeTask2Request {
	s.ClientId = &v
	return s
}

func (s *SelectComputeTask2Request) Validate() error {
	return dara.Validate(s)
}
