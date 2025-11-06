// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForTransferOutByAuthorizationCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTransferOutParamList(v []*SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList) *SaveBatchTaskForTransferOutByAuthorizationCodeRequest
	GetTransferOutParamList() []*SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList
}

type SaveBatchTaskForTransferOutByAuthorizationCodeRequest struct {
	// This parameter is required.
	TransferOutParamList []*SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList `json:"TransferOutParamList,omitempty" xml:"TransferOutParamList,omitempty" type:"Repeated"`
}

func (s SaveBatchTaskForTransferOutByAuthorizationCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForTransferOutByAuthorizationCodeRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeRequest) GetTransferOutParamList() []*SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList {
	return s.TransferOutParamList
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeRequest) SetTransferOutParamList(v []*SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList) *SaveBatchTaskForTransferOutByAuthorizationCodeRequest {
	s.TransferOutParamList = v
	return s
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeRequest) Validate() error {
	if s.TransferOutParamList != nil {
		for _, item := range s.TransferOutParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList struct {
	// example:
	//
	// Test2o#Lck
	AuthorizationCode *string `json:"AuthorizationCode,omitempty" xml:"AuthorizationCode,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList) GetAuthorizationCode() *string {
	return s.AuthorizationCode
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList) SetAuthorizationCode(v string) *SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList {
	s.AuthorizationCode = &v
	return s
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList) SetDomainName(v string) *SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList {
	s.DomainName = &v
	return s
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeRequestTransferOutParamList) Validate() error {
	return dara.Validate(s)
}
