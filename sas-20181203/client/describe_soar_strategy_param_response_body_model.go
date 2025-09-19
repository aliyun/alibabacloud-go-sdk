// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategyParamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v string) *DescribeSoarStrategyParamResponseBody
	GetParams() *string
	SetProcessInfo(v string) *DescribeSoarStrategyParamResponseBody
	GetProcessInfo() *string
	SetRequestId(v string) *DescribeSoarStrategyParamResponseBody
	GetRequestId() *string
}

type DescribeSoarStrategyParamResponseBody struct {
	// The parameters of the policy.
	//
	// example:
	//
	// {"summary":[{"name":"email","type":"String","isRequired":false,"fromProperty":"notifyConfig.email"}]}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The process information of the policy.
	//
	// example:
	//
	// {"edges":[{"level":0,"removeFlag":0,"source":1,"target":8}]}
	ProcessInfo *string `json:"ProcessInfo,omitempty" xml:"ProcessInfo,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6673D49C-A9AB-40DD-B4A2-B92306701AE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSoarStrategyParamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyParamResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyParamResponseBody) GetParams() *string {
	return s.Params
}

func (s *DescribeSoarStrategyParamResponseBody) GetProcessInfo() *string {
	return s.ProcessInfo
}

func (s *DescribeSoarStrategyParamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSoarStrategyParamResponseBody) SetParams(v string) *DescribeSoarStrategyParamResponseBody {
	s.Params = &v
	return s
}

func (s *DescribeSoarStrategyParamResponseBody) SetProcessInfo(v string) *DescribeSoarStrategyParamResponseBody {
	s.ProcessInfo = &v
	return s
}

func (s *DescribeSoarStrategyParamResponseBody) SetRequestId(v string) *DescribeSoarStrategyParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSoarStrategyParamResponseBody) Validate() error {
	return dara.Validate(s)
}
