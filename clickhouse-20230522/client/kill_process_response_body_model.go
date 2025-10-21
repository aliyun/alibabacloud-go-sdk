// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *KillProcessResponseBodyData) *KillProcessResponseBody
	GetData() *KillProcessResponseBodyData
	SetRequestId(v string) *KillProcessResponseBody
	GetRequestId() *string
}

type KillProcessResponseBody struct {
	// The data returned.
	Data *KillProcessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KillProcessResponseBody) GoString() string {
	return s.String()
}

func (s *KillProcessResponseBody) GetData() *KillProcessResponseBodyData {
	return s.Data
}

func (s *KillProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KillProcessResponseBody) SetData(v *KillProcessResponseBodyData) *KillProcessResponseBody {
	s.Data = v
	return s
}

func (s *KillProcessResponseBody) SetRequestId(v string) *KillProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *KillProcessResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KillProcessResponseBodyData struct {
	// The number of queries that are terminated.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s KillProcessResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s KillProcessResponseBodyData) GoString() string {
	return s.String()
}

func (s *KillProcessResponseBodyData) GetCount() *int64 {
	return s.Count
}

func (s *KillProcessResponseBodyData) GetDBInstanceID() *int32 {
	return s.DBInstanceID
}

func (s *KillProcessResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *KillProcessResponseBodyData) SetCount(v int64) *KillProcessResponseBodyData {
	s.Count = &v
	return s
}

func (s *KillProcessResponseBodyData) SetDBInstanceID(v int32) *KillProcessResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *KillProcessResponseBodyData) SetDBInstanceName(v string) *KillProcessResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *KillProcessResponseBodyData) Validate() error {
	return dara.Validate(s)
}
