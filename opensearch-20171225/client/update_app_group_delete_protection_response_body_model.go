// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppGroupDeleteProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAppGroupDeleteProtectionResponseBody
	GetRequestId() *string
	SetResult(v *UpdateAppGroupDeleteProtectionResponseBodyResult) *UpdateAppGroupDeleteProtectionResponseBody
	GetResult() *UpdateAppGroupDeleteProtectionResponseBodyResult
}

type UpdateAppGroupDeleteProtectionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 80326EFE-485F-46E7-B291-5A1DD08D2198
	RequestId *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateAppGroupDeleteProtectionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateAppGroupDeleteProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppGroupDeleteProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppGroupDeleteProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppGroupDeleteProtectionResponseBody) GetResult() *UpdateAppGroupDeleteProtectionResponseBodyResult {
	return s.Result
}

func (s *UpdateAppGroupDeleteProtectionResponseBody) SetRequestId(v string) *UpdateAppGroupDeleteProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppGroupDeleteProtectionResponseBody) SetResult(v *UpdateAppGroupDeleteProtectionResponseBodyResult) *UpdateAppGroupDeleteProtectionResponseBody {
	s.Result = v
	return s
}

func (s *UpdateAppGroupDeleteProtectionResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAppGroupDeleteProtectionResponseBodyResult struct {
	// example:
	//
	// ops-cn-m7r1ywo2h1b
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s UpdateAppGroupDeleteProtectionResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppGroupDeleteProtectionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateAppGroupDeleteProtectionResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAppGroupDeleteProtectionResponseBodyResult) SetInstanceId(v string) *UpdateAppGroupDeleteProtectionResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *UpdateAppGroupDeleteProtectionResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
