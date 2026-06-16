// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InitializeV2ResponseBody
	GetCode() *string
	SetMessage(v string) *InitializeV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitializeV2ResponseBody
	GetRequestId() *string
	SetResult(v *InitializeV2ResponseBodyResult) *InitializeV2ResponseBody
	GetResult() *InitializeV2ResponseBodyResult
}

type InitializeV2ResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86C40EC3-5940-5F47-995C-BFE90B70E540
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *InitializeV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s InitializeV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitializeV2ResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitializeV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitializeV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitializeV2ResponseBody) GetResult() *InitializeV2ResponseBodyResult {
	return s.Result
}

func (s *InitializeV2ResponseBody) SetCode(v string) *InitializeV2ResponseBody {
	s.Code = &v
	return s
}

func (s *InitializeV2ResponseBody) SetMessage(v string) *InitializeV2ResponseBody {
	s.Message = &v
	return s
}

func (s *InitializeV2ResponseBody) SetRequestId(v string) *InitializeV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeV2ResponseBody) SetResult(v *InitializeV2ResponseBodyResult) *InitializeV2ResponseBody {
	s.Result = v
	return s
}

func (s *InitializeV2ResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitializeV2ResponseBodyResult struct {
	// example:
	//
	// ***
	ClientCfg *string `json:"ClientCfg,omitempty" xml:"ClientCfg,omitempty"`
	// example:
	//
	// hksb7ba1b28130d24e015d*********
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
	// example:
	//
	// http****
	TransactionUrl *string `json:"TransactionUrl,omitempty" xml:"TransactionUrl,omitempty"`
}

func (s InitializeV2ResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s InitializeV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InitializeV2ResponseBodyResult) GetClientCfg() *string {
	return s.ClientCfg
}

func (s *InitializeV2ResponseBodyResult) GetProtocol() *string {
	return s.Protocol
}

func (s *InitializeV2ResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *InitializeV2ResponseBodyResult) GetTransactionUrl() *string {
	return s.TransactionUrl
}

func (s *InitializeV2ResponseBodyResult) SetClientCfg(v string) *InitializeV2ResponseBodyResult {
	s.ClientCfg = &v
	return s
}

func (s *InitializeV2ResponseBodyResult) SetProtocol(v string) *InitializeV2ResponseBodyResult {
	s.Protocol = &v
	return s
}

func (s *InitializeV2ResponseBodyResult) SetTransactionId(v string) *InitializeV2ResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *InitializeV2ResponseBodyResult) SetTransactionUrl(v string) *InitializeV2ResponseBodyResult {
	s.TransactionUrl = &v
	return s
}

func (s *InitializeV2ResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
