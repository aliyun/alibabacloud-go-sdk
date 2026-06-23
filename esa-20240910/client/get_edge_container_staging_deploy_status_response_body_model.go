// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerStagingDeployStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContainersReady(v string) *GetEdgeContainerStagingDeployStatusResponseBody
	GetContainersReady() *string
	SetCreationTimestamp(v string) *GetEdgeContainerStagingDeployStatusResponseBody
	GetCreationTimestamp() *string
	SetInitialized(v string) *GetEdgeContainerStagingDeployStatusResponseBody
	GetInitialized() *string
	SetPhase(v string) *GetEdgeContainerStagingDeployStatusResponseBody
	GetPhase() *string
	SetPodRestartState(v *GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState) *GetEdgeContainerStagingDeployStatusResponseBody
	GetPodRestartState() *GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState
	SetReady(v string) *GetEdgeContainerStagingDeployStatusResponseBody
	GetReady() *string
	SetRequestId(v string) *GetEdgeContainerStagingDeployStatusResponseBody
	GetRequestId() *string
	SetScheduled(v string) *GetEdgeContainerStagingDeployStatusResponseBody
	GetScheduled() *string
	SetVIPs(v []*string) *GetEdgeContainerStagingDeployStatusResponseBody
	GetVIPs() []*string
}

type GetEdgeContainerStagingDeployStatusResponseBody struct {
	// Indicates whether the container status is ready. Valid values:
	//
	// - **ok**: Ready.
	//
	// - **unready**: Not ready.
	//
	// example:
	//
	// ok
	ContainersReady *string `json:"ContainersReady,omitempty" xml:"ContainersReady,omitempty"`
	// The creation time (UNIX timestamp).
	//
	// example:
	//
	// 2024-09-24T06:46:35Z
	CreationTimestamp *string `json:"CreationTimestamp,omitempty" xml:"CreationTimestamp,omitempty"`
	// The container initialization status. Valid values:
	//
	// - **ok**: Succeeded.
	//
	// - **unready**: Not completed.
	//
	// example:
	//
	// ok
	Initialized *string `json:"Initialized,omitempty" xml:"Initialized,omitempty"`
	// The status of the container in the staging environment. Valid values:
	//
	// - NoContainer: no container.
	//
	// - Pending: pending deployment.
	//
	// - ContainerCreating: the container is being created.
	//
	// - Running: running.
	//
	// - Succeeded: completed.
	//
	// - ImagePullBackOff: image pull failed.
	//
	// - CrashLoopBackOff: abnormal container startup.
	//
	// - Failed: failed.
	//
	// - Unknown: unknown.
	//
	// example:
	//
	// Running
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The container restart status.
	PodRestartState *GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState `json:"PodRestartState,omitempty" xml:"PodRestartState,omitempty" type:"Struct"`
	// Indicates whether the container is ready to receive traffic. Valid values:
	//
	// - **ok**: Ready.
	//
	// - **unready**: Not ready.
	//
	// example:
	//
	// ok
	Ready *string `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2F2C992B-3FE2-5EBB-A61F-F9DD4EB257DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The container scheduling status. Valid values:
	//
	// - **ok**: Succeeded.
	//
	// - **unready**: Not completed.
	//
	// example:
	//
	// ok
	Scheduled *string `json:"Scheduled,omitempty" xml:"Scheduled,omitempty"`
	// The list of VIPs.
	VIPs []*string `json:"VIPs,omitempty" xml:"VIPs,omitempty" type:"Repeated"`
}

func (s GetEdgeContainerStagingDeployStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerStagingDeployStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) GetContainersReady() *string {
	return s.ContainersReady
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) GetCreationTimestamp() *string {
	return s.CreationTimestamp
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) GetInitialized() *string {
	return s.Initialized
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) GetPhase() *string {
	return s.Phase
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) GetPodRestartState() *GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState {
	return s.PodRestartState
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) GetReady() *string {
	return s.Ready
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) GetScheduled() *string {
	return s.Scheduled
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) GetVIPs() []*string {
	return s.VIPs
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) SetContainersReady(v string) *GetEdgeContainerStagingDeployStatusResponseBody {
	s.ContainersReady = &v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) SetCreationTimestamp(v string) *GetEdgeContainerStagingDeployStatusResponseBody {
	s.CreationTimestamp = &v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) SetInitialized(v string) *GetEdgeContainerStagingDeployStatusResponseBody {
	s.Initialized = &v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) SetPhase(v string) *GetEdgeContainerStagingDeployStatusResponseBody {
	s.Phase = &v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) SetPodRestartState(v *GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState) *GetEdgeContainerStagingDeployStatusResponseBody {
	s.PodRestartState = v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) SetReady(v string) *GetEdgeContainerStagingDeployStatusResponseBody {
	s.Ready = &v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) SetRequestId(v string) *GetEdgeContainerStagingDeployStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) SetScheduled(v string) *GetEdgeContainerStagingDeployStatusResponseBody {
	s.Scheduled = &v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) SetVIPs(v []*string) *GetEdgeContainerStagingDeployStatusResponseBody {
	s.VIPs = v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusResponseBody) Validate() error {
	if s.PodRestartState != nil {
		if err := s.PodRestartState.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState struct {
	// The reason for the last restart.
	//
	// example:
	//
	// OOMKilled
	LastTerminatedReason *string `json:"LastTerminatedReason,omitempty" xml:"LastTerminatedReason,omitempty"`
	// The number of restarts.
	//
	// example:
	//
	// 1
	RestartCount *int32 `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
}

func (s GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState) GetLastTerminatedReason() *string {
	return s.LastTerminatedReason
}

func (s *GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState) GetRestartCount() *int32 {
	return s.RestartCount
}

func (s *GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState) SetLastTerminatedReason(v string) *GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState {
	s.LastTerminatedReason = &v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState) SetRestartCount(v int32) *GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState {
	s.RestartCount = &v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusResponseBodyPodRestartState) Validate() error {
	return dara.Validate(s)
}
