// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAgentClientInstallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAegisCelintInstallResposeList(v []*OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList) *OperateAgentClientInstallResponseBody
	GetAegisCelintInstallResposeList() []*OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList
	SetRequestId(v string) *OperateAgentClientInstallResponseBody
	GetRequestId() *string
}

type OperateAgentClientInstallResponseBody struct {
	// An array that consists of the returned results.
	AegisCelintInstallResposeList []*OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList `json:"AegisCelintInstallResposeList,omitempty" xml:"AegisCelintInstallResposeList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// AE79B457-877C-51C6-AD72-0D34A025D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateAgentClientInstallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateAgentClientInstallResponseBody) GoString() string {
	return s.String()
}

func (s *OperateAgentClientInstallResponseBody) GetAegisCelintInstallResposeList() []*OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList {
	return s.AegisCelintInstallResposeList
}

func (s *OperateAgentClientInstallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateAgentClientInstallResponseBody) SetAegisCelintInstallResposeList(v []*OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList) *OperateAgentClientInstallResponseBody {
	s.AegisCelintInstallResposeList = v
	return s
}

func (s *OperateAgentClientInstallResponseBody) SetRequestId(v string) *OperateAgentClientInstallResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateAgentClientInstallResponseBody) Validate() error {
	if s.AegisCelintInstallResposeList != nil {
		for _, item := range s.AegisCelintInstallResposeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList struct {
	// The ID of the server.
	//
	// example:
	//
	// i-uf6j8vq9l4r5ntht****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the installation task.
	//
	// example:
	//
	// 2856
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 1587bedb-fdb4-48c4-9330-****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList) String() string {
	return dara.Prettify(s)
}

func (s OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList) GoString() string {
	return s.String()
}

func (s *OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList) GetRecordId() *int64 {
	return s.RecordId
}

func (s *OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList) GetUuid() *string {
	return s.Uuid
}

func (s *OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList) SetInstanceId(v string) *OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList {
	s.InstanceId = &v
	return s
}

func (s *OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList) SetRecordId(v int64) *OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList {
	s.RecordId = &v
	return s
}

func (s *OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList) SetUuid(v string) *OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList {
	s.Uuid = &v
	return s
}

func (s *OperateAgentClientInstallResponseBodyAegisCelintInstallResposeList) Validate() error {
	return dara.Validate(s)
}
