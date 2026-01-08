// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAppResponseBody
	GetRequestId() *string
	SetResult(v *UpdateAppResponseBodyResult) *UpdateAppResponseBody
	GetResult() *UpdateAppResponseBodyResult
}

type UpdateAppResponseBody struct {
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppResponseBody) GetResult() *UpdateAppResponseBodyResult {
	return s.Result
}

func (s *UpdateAppResponseBody) SetRequestId(v string) *UpdateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppResponseBody) SetResult(v *UpdateAppResponseBodyResult) *UpdateAppResponseBody {
	s.Result = v
	return s
}

func (s *UpdateAppResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAppResponseBodyResult struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s UpdateAppResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAppResponseBodyResult) SetInstanceId(v string) *UpdateAppResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *UpdateAppResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
