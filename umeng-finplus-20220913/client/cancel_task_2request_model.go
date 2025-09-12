// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTask2Request interface {
	dara.Model
	String() string
	GoString() string
	SetBcId(v int64) *CancelTask2Request
	GetBcId() *int64
	SetClientId(v int64) *CancelTask2Request
	GetClientId() *int64
}

type CancelTask2Request struct {
	BcId     *int64 `json:"bcId,omitempty" xml:"bcId,omitempty"`
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
}

func (s CancelTask2Request) String() string {
	return dara.Prettify(s)
}

func (s CancelTask2Request) GoString() string {
	return s.String()
}

func (s *CancelTask2Request) GetBcId() *int64 {
	return s.BcId
}

func (s *CancelTask2Request) GetClientId() *int64 {
	return s.ClientId
}

func (s *CancelTask2Request) SetBcId(v int64) *CancelTask2Request {
	s.BcId = &v
	return s
}

func (s *CancelTask2Request) SetClientId(v int64) *CancelTask2Request {
	s.ClientId = &v
	return s
}

func (s *CancelTask2Request) Validate() error {
	return dara.Validate(s)
}
