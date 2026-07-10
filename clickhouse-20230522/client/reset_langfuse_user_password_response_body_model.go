// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetLangfuseUserPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ResetLangfuseUserPasswordResponseBodyData) *ResetLangfuseUserPasswordResponseBody
	GetData() *ResetLangfuseUserPasswordResponseBodyData
	SetRequestId(v string) *ResetLangfuseUserPasswordResponseBody
	GetRequestId() *string
}

type ResetLangfuseUserPasswordResponseBody struct {
	// The response result.
	Data *ResetLangfuseUserPasswordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetLangfuseUserPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetLangfuseUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetLangfuseUserPasswordResponseBody) GetData() *ResetLangfuseUserPasswordResponseBodyData {
	return s.Data
}

func (s *ResetLangfuseUserPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetLangfuseUserPasswordResponseBody) SetData(v *ResetLangfuseUserPasswordResponseBodyData) *ResetLangfuseUserPasswordResponseBody {
	s.Data = v
	return s
}

func (s *ResetLangfuseUserPasswordResponseBody) SetRequestId(v string) *ResetLangfuseUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetLangfuseUserPasswordResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetLangfuseUserPasswordResponseBodyData struct {
	// The email address of the user.
	//
	// example:
	//
	// john@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s ResetLangfuseUserPasswordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResetLangfuseUserPasswordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetLangfuseUserPasswordResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *ResetLangfuseUserPasswordResponseBodyData) SetEmail(v string) *ResetLangfuseUserPasswordResponseBodyData {
	s.Email = &v
	return s
}

func (s *ResetLangfuseUserPasswordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
