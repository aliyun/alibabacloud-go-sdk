// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeResourceAuthUserMappingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetComputeResourceAuthUserMappingsResponseBodyData) *GetComputeResourceAuthUserMappingsResponseBody
	GetData() *GetComputeResourceAuthUserMappingsResponseBodyData
	SetRequestId(v string) *GetComputeResourceAuthUserMappingsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetComputeResourceAuthUserMappingsResponseBody
	GetSuccess() *bool
}

type GetComputeResourceAuthUserMappingsResponseBody struct {
	// The returned data.
	Data *GetComputeResourceAuthUserMappingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 0000-ABCD-EF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetComputeResourceAuthUserMappingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetComputeResourceAuthUserMappingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeResourceAuthUserMappingsResponseBody) GetData() *GetComputeResourceAuthUserMappingsResponseBodyData {
	return s.Data
}

func (s *GetComputeResourceAuthUserMappingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetComputeResourceAuthUserMappingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetComputeResourceAuthUserMappingsResponseBody) SetData(v *GetComputeResourceAuthUserMappingsResponseBodyData) *GetComputeResourceAuthUserMappingsResponseBody {
	s.Data = v
	return s
}

func (s *GetComputeResourceAuthUserMappingsResponseBody) SetRequestId(v string) *GetComputeResourceAuthUserMappingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComputeResourceAuthUserMappingsResponseBody) SetSuccess(v bool) *GetComputeResourceAuthUserMappingsResponseBody {
	s.Success = &v
	return s
}

func (s *GetComputeResourceAuthUserMappingsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetComputeResourceAuthUserMappingsResponseBodyData struct {
	// The list of mapped account information.
	Accounts []*GetComputeResourceAuthUserMappingsResponseBodyDataAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// The authentication type, such as LDAP.
	//
	// example:
	//
	// ldap
	HadoopAuthType *string `json:"HadoopAuthType,omitempty" xml:"HadoopAuthType,omitempty"`
}

func (s GetComputeResourceAuthUserMappingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetComputeResourceAuthUserMappingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetComputeResourceAuthUserMappingsResponseBodyData) GetAccounts() []*GetComputeResourceAuthUserMappingsResponseBodyDataAccounts {
	return s.Accounts
}

func (s *GetComputeResourceAuthUserMappingsResponseBodyData) GetHadoopAuthType() *string {
	return s.HadoopAuthType
}

func (s *GetComputeResourceAuthUserMappingsResponseBodyData) SetAccounts(v []*GetComputeResourceAuthUserMappingsResponseBodyDataAccounts) *GetComputeResourceAuthUserMappingsResponseBodyData {
	s.Accounts = v
	return s
}

func (s *GetComputeResourceAuthUserMappingsResponseBodyData) SetHadoopAuthType(v string) *GetComputeResourceAuthUserMappingsResponseBodyData {
	s.HadoopAuthType = &v
	return s
}

func (s *GetComputeResourceAuthUserMappingsResponseBodyData) Validate() error {
	if s.Accounts != nil {
		for _, item := range s.Accounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetComputeResourceAuthUserMappingsResponseBodyDataAccounts struct {
	// The Alibaba Cloud UID.
	//
	// example:
	//
	// 12747300953xxx62
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username in the target system, such as an LDAP account.
	//
	// example:
	//
	// jsmitxxxx
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetComputeResourceAuthUserMappingsResponseBodyDataAccounts) String() string {
	return dara.Prettify(s)
}

func (s GetComputeResourceAuthUserMappingsResponseBodyDataAccounts) GoString() string {
	return s.String()
}

func (s *GetComputeResourceAuthUserMappingsResponseBodyDataAccounts) GetUserId() *string {
	return s.UserId
}

func (s *GetComputeResourceAuthUserMappingsResponseBodyDataAccounts) GetUsername() *string {
	return s.Username
}

func (s *GetComputeResourceAuthUserMappingsResponseBodyDataAccounts) SetUserId(v string) *GetComputeResourceAuthUserMappingsResponseBodyDataAccounts {
	s.UserId = &v
	return s
}

func (s *GetComputeResourceAuthUserMappingsResponseBodyDataAccounts) SetUsername(v string) *GetComputeResourceAuthUserMappingsResponseBodyDataAccounts {
	s.Username = &v
	return s
}

func (s *GetComputeResourceAuthUserMappingsResponseBodyDataAccounts) Validate() error {
	return dara.Validate(s)
}
