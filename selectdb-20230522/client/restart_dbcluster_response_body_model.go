// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RestartDBClusterResponseBodyData) *RestartDBClusterResponseBody
	GetData() *RestartDBClusterResponseBodyData
	SetRequestId(v string) *RestartDBClusterResponseBody
	GetRequestId() *string
}

type RestartDBClusterResponseBody struct {
	// The information returned.
	Data *RestartDBClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BD0D0B17-C145-5B91-BFC2-6791927EE973
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDBClusterResponseBody) GetData() *RestartDBClusterResponseBodyData {
	return s.Data
}

func (s *RestartDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartDBClusterResponseBody) SetData(v *RestartDBClusterResponseBodyData) *RestartDBClusterResponseBody {
	s.Data = v
	return s
}

func (s *RestartDBClusterResponseBody) SetRequestId(v string) *RestartDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type RestartDBClusterResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// selectdb-cn-7213c8y****-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s RestartDBClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RestartDBClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestartDBClusterResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RestartDBClusterResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RestartDBClusterResponseBodyData) SetDBClusterId(v string) *RestartDBClusterResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *RestartDBClusterResponseBodyData) SetDBInstanceId(v string) *RestartDBClusterResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *RestartDBClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
